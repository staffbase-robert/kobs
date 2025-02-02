package harbor

import (
	"net/http"

	"github.com/kobsio/kobs/pkg/api/clusters"
	"github.com/kobsio/kobs/pkg/api/middleware/errresponse"
	"github.com/kobsio/kobs/pkg/api/plugins/plugin"
	"github.com/kobsio/kobs/plugins/harbor/pkg/instance"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Route is the route under which the plugin should be registered in our router for the rest api.
const (
	Route = "/harbor"
)

var (
	log = logrus.WithFields(logrus.Fields{"package": "harbor"})
)

// Config is the structure of the configuration for the Harbor plugin.
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

func (router *Router) getProjects(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")

	log.WithFields(logrus.Fields{"name": name, "page": page, "pageSize": pageSize}).Tracef("getProjects")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	projectsData, err := i.GetProjects(r.Context(), page, pageSize)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get projects")
		return
	}

	render.JSON(w, r, projectsData)
}

func (router *Router) getRepositories(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	projectName := r.URL.Query().Get("projectName")
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")

	log.WithFields(logrus.Fields{"name": name, "projectName": projectName, "query": query, "page": page, "pageSize": pageSize}).Tracef("getRepositories")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	repositoriesData, err := i.GetRepositories(r.Context(), projectName, query, page, pageSize)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get repositories")
		return
	}

	render.JSON(w, r, repositoriesData)
}

func (router *Router) getArtifacts(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	projectName := r.URL.Query().Get("projectName")
	repositoryName := r.URL.Query().Get("repositoryName")
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")

	log.WithFields(logrus.Fields{"name": name, "projectName": projectName, "repositoryName": repositoryName, "query": query, "page": page, "pageSize": pageSize}).Tracef("getArtifacts")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	artifactsData, err := i.GetArtifacts(r.Context(), projectName, repositoryName, query, page, pageSize)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get artifacts")
		return
	}

	render.JSON(w, r, artifactsData)
}

func (router *Router) getArtifact(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	projectName := r.URL.Query().Get("projectName")
	repositoryName := r.URL.Query().Get("repositoryName")
	artifactReference := r.URL.Query().Get("artifactReference")

	log.WithFields(logrus.Fields{"name": name, "projectName": projectName, "repositoryName": repositoryName, "artifactReference": artifactReference}).Tracef("getArtifact")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	artifact, err := i.GetArtifact(r.Context(), projectName, repositoryName, artifactReference)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get artifact")
		return
	}

	render.JSON(w, r, artifact)
}

func (router *Router) getVulnerabilities(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	projectName := r.URL.Query().Get("projectName")
	repositoryName := r.URL.Query().Get("repositoryName")
	artifactReference := r.URL.Query().Get("artifactReference")

	log.WithFields(logrus.Fields{"name": name, "projectName": projectName, "repositoryName": repositoryName, "artifactReference": artifactReference}).Tracef("getVulnerabilities")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	vulnerabilities, err := i.GetVulnerabilities(r.Context(), projectName, repositoryName, artifactReference)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get vulnerabilities")
		return
	}

	render.JSON(w, r, vulnerabilities)
}

func (router *Router) getBuildHistory(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	projectName := r.URL.Query().Get("projectName")
	repositoryName := r.URL.Query().Get("repositoryName")
	artifactReference := r.URL.Query().Get("artifactReference")

	log.WithFields(logrus.Fields{"name": name, "projectName": projectName, "repositoryName": repositoryName, "artifactReference": artifactReference}).Tracef("getBuildHistory")

	i := router.getInstance(name)
	if i == nil {
		errresponse.Render(w, r, nil, http.StatusBadRequest, "Could not find instance name")
		return
	}

	buildHistory, err := i.GetBuildHistory(r.Context(), projectName, repositoryName, artifactReference)
	if err != nil {
		errresponse.Render(w, r, err, http.StatusInternalServerError, "Could not get build history")
		return
	}

	render.JSON(w, r, buildHistory)
}

// Register returns a new router which can be used in the router for the kobs rest api.
func Register(clusters *clusters.Clusters, plugins *plugin.Plugins, config Config) chi.Router {
	var instances []*instance.Instance

	for _, cfg := range config {
		instance, err := instance.New(cfg)
		if err != nil {
			log.WithError(err).WithFields(logrus.Fields{"name": cfg.Name}).Fatalf("Could not create Harbor instance")
		}

		instances = append(instances, instance)

		var options map[string]interface{}
		options = make(map[string]interface{})
		options["address"] = cfg.Address

		plugins.Append(plugin.Plugin{
			Name:        cfg.Name,
			DisplayName: cfg.DisplayName,
			Description: cfg.Description,
			Type:        "harbor",
			Options:     options,
		})
	}

	router := Router{
		chi.NewRouter(),
		clusters,
		instances,
	}

	router.Get("/projects/{name}", router.getProjects)
	router.Get("/repositories/{name}", router.getRepositories)
	router.Get("/artifacts/{name}", router.getArtifacts)
	router.Get("/artifact/{name}", router.getArtifact)
	router.Get("/vulnerabilities/{name}", router.getVulnerabilities)
	router.Get("/buildhistory/{name}", router.getBuildHistory)

	return router
}
