package okctl

import (
	"ob-operator-mcp-server/pkg/toolsets"
)

func InitToolsets(passedToolSets []string) (*toolsets.ToolSetsGroup, error) {
	// init toolSets group
	tsg := toolsets.NewToolSetGroup()
	obclusters := toolsets.NewToolSet("obclusters", "oceanbase cluster related tools").
		AddReadTools(
			toolsets.NewServerTool(listAllClusters()),
			toolsets.NewServerTool(showCluster()),
		).
		AddWriteTools(
			toolsets.NewServerTool(createCluster()),
			toolsets.NewServerTool(deleteCluster()),
			toolsets.NewServerTool(updateCluster()),
			toolsets.NewServerTool(upgradeCluster()),
			toolsets.NewServerTool(scaleCluster()),
		)
	obtenants := toolsets.NewToolSet("obtenants", "oceanbase tenant related tools").AddReadTools().AddWriteTools()
	obbackupproxies := toolsets.NewToolSet("obbackupproxy", "oceanbase backup proxy related tools").AddReadTools().AddWriteTools()
	installs := toolsets.NewToolSet("installs", "oceanbase install related tools").AddReadTools().AddWriteTools()
	updates := toolsets.NewToolSet("updates", "oceanbase update related tools").AddReadTools().AddWriteTools()

	// Add toolsets to the group
	tsg.AddToolSet(obclusters)
	tsg.AddToolSet(obtenants)
	tsg.AddToolSet(obbackupproxies)
	tsg.AddToolSet(installs)
	tsg.AddToolSet(updates)

	// enable passed toolSets
	if err := tsg.EnableToolSets(passedToolSets); err != nil {
		return nil, err
	}

	return tsg, nil
}
