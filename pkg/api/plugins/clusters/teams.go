package clusters

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Team struct {
	Cluster      string
	Namespace    string
	Name         string
	Description  string
	Logo         string
	Applications []Application
}

type Application struct {
	Cluster   string
	Namespace string
	Name      string
}

// generateTeams is used to generate a list of teams. This list contains the cluster, namespace and name where we can
// found the Team CR. It also contains the description and logo for the team, because we need them for the overview
// page. Last but not least it contains a list of applications, with the cluster, namespace and name, so that we can
// retrieve the complete Application CR on request.
// To generate this list we get all teams and applications from all clusters. Then we add the team to this list, when
// the team with the same name doesn't already exists. This means the name of the Team CR must be unique across clusters
// and namespaces. Finally we loop through all applications and check if the application contains the team name in the
// teams property. If this is the case we add the application to the team.
func (c *Clusters) generateTeams() {
	sleep, err := time.ParseDuration(cacheDurationTeams)
	if err != nil {
		log.WithError(err).Errorf("Invalide cache duration for teams, use default cache duration of 60m")
		sleep = time.Duration(60 * time.Minute)
	}

	for {
		time.Sleep(60 * time.Second)
		log.Infof("Generate teams")
		ctx := context.Background()

		var cachedTeams []Team

		for _, c := range c.clusters {
			clusterName := c.GetName()

			teams, err := c.GetTeams(ctx)
			if err != nil {
				log.WithError(err).WithFields(logrus.Fields{"cluster": clusterName}).Errorf("Could not get teams")
				continue
			}

			applications, err := c.GetApplications(ctx, "")
			if err != nil {
				log.WithError(err).WithFields(logrus.Fields{"cluster": clusterName}).Errorf("Could not get applications")
				continue
			}

			for _, team := range teams {
				// Skip the team if it already exists, because teams must be unique accross clusters and namespaces.
				if doesTeamExists(cachedTeams, team.Name) {
					continue
				}

				var teamApplications []Application
				for _, application := range applications {
					if containsTeam(application.Teams, team.Name) {
						teamApplications = append(teamApplications, Application{
							Cluster:   application.Cluster,
							Namespace: application.Namespace,
							Name:      application.Name,
						})
					}
				}

				cachedTeams = append(cachedTeams, Team{
					Cluster:      team.Cluster,
					Namespace:    team.Namespace,
					Name:         team.Name,
					Description:  team.Description,
					Logo:         team.Logo,
					Applications: teamApplications,
				})
			}
		}

		c.teams = cachedTeams

		time.Sleep(sleep)
	}
}

// doesTeamExists checks if the given team name exists in a slice of teams.
func doesTeamExists(teams []Team, name string) bool {
	for _, team := range teams {
		if team.Name == name {
			return true
		}
	}

	return false
}

// containsTeam checks if the given team name exists in a slice of teams.
func containsTeam(teams []string, team string) bool {
	for _, t := range teams {
		if t == team {
			return true
		}
	}

	return false
}

// getTeamData returns the team for the given team name.
func getTeamData(teams []Team, name string) *Team {
	for _, team := range teams {
		if team.Name == name {
			return &team
		}
	}

	return nil
}
