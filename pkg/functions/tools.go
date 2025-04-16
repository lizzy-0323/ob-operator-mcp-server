package functions

import (
	"ob-operator-mcp-server/pkg/toolsets"
)

func InitToolsets() *toolsets.ToolSetsGroup {
	// init toolsets group
	tsg := toolsets.NewToolSetGroup()
	obclusters := toolsets.NewToolSet("obclusters", "oceanbase cluster related tools").AddReadTools(
		toolsets.NewServerTool(listAllClusters()),
	).AddWriteTools(
		toolsets.NewServerTool(createCluster()),
		toolsets.NewServerTool(deleteCluster()),
	)
	obtenants := toolsets.NewToolSet("obtenants", "oceanbase tenant related tools").AddReadTools().AddWriteTools()
	obbackupproxies := toolsets.NewToolSet("obbackupproxies", "oceanbase backup proxy related tools").AddReadTools().AddWriteTools()
	installs := toolsets.NewToolSet("installs", "oceanbase install related tools").AddReadTools().AddWriteTools()
	updates := toolsets.NewToolSet("updates", "oceanbase update related tools").AddReadTools().AddWriteTools()

	tsg.AddToolSet(obclusters)
	tsg.AddToolSet(obtenants)
	tsg.AddToolSet(obbackupproxies)
	tsg.AddToolSet(installs)
	tsg.AddToolSet(updates)

	return tsg
}
