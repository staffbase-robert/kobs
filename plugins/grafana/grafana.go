package grafana

import (
	"net/http"

	"github.com/kobsio/kobs/pkg/api/clusters"
	"github.com/kobsio/kobs/pkg/api/middleware/errresponse"
	"github.com/kobsio/kobs/pkg/api/plugins/plugin"
	"github.com/kobsio/kobs/plugins/grafana/pkg/instance"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Route is the route under which the plugin should be registered in our router for the rest api.
const (
	Route = "/grafana"
)

var (
	log = logrus.WithFields(logrus.Fields{"package": "grafana"})
)

// Config is the structure of the configuration for the grafana plugin.
type Config []instance.Config

// Router implements the router for the resources plugin, which can be registered in the router for our rest api.
type Router struct {
	*chi.Mux
	clusters  *clusters.Clusters
	instances []*instance.Instance
}

func (router *Router) getInstance(name string) *instance.Instance {
	for _, i := range router.instances {
		if i.Name == name {
			return i
		}
	}

	return nil
}

func (router *Router) getDashboards(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	query := r.URL.Query().Get("query")
	uids := r.URL.Query()["uid"]

	log.WithFields(logrus.Fields{"name": name, "query": query, "uids": uids}).Tracef("getDashboards")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	if uids != nil {
		var dashboards []instance.Dashboard
		for _, uid := range uids {
			dashboard, err := i.GetDashboard(r.Context(), uid)
			if err != nil {
				errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get dashboard")
				return
			}

			dashboards = append(dashboards, *dashboard)
		}

		render.JSON(w, r, dashboards)
		return
	}

	dashboards, err := i.GetDashboards(r.Context(), query)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get dashboards")
		return
	}

	render.JSON(w, r, dashboards)
}

// Register returns a new router which can be used in the router for the kobs rest api.
func Register(clusters *clusters.Clusters, plugins *plugin.Plugins, config Config) chi.Router {
	var instances []*instance.Instance

	for _, cfg := range config {
		instance, err := instance.New(cfg)
		if err != nil {
			log.WithError(err).WithFields(logrus.Fields{"name": cfg.Name}).Fatalf("Could not create Grafana instance")
		}

		instances = append(instances, instance)

		var options map[string]interface{}
		options = make(map[string]interface{})
		options["internalAddress"] = cfg.InternalAddress
		options["publicAddress"] = cfg.PublicAddress

		plugins.Append(plugin.Plugin{
			Name:        cfg.Name,
			DisplayName: cfg.DisplayName,
			Description: cfg.Description,
			Type:        "grafana",
			Options:     options,
		})
	}

	router := Router{
		chi.NewRouter(),
		clusters,
		instances,
	}

	router.Get("/dashboards/{name}", router.getDashboards)

	return router
}
