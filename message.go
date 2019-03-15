package main

// GroupResultToArtifacts func
func GroupResultToArtifacts(messages []Result) []Artifact {
	artifacts := []Artifact{}
	errorGroups := make(map[string][]Result)

	for _, m := range messages {
		if g, ok := errorGroups[string(m.IntegrationArtifact.Name)]; ok {
			errorGroups[string(m.IntegrationArtifact.Name)] = append(g, m)
		} else {
			errorGroups[string(m.IntegrationArtifact.Name)] = []Result{m}
		}
	}

	for k, v := range errorGroups {
		artifacts = append(artifacts, Artifact{
			ArtifactName: k,
			Errors:       v,
		})
	}
	return artifacts
}
