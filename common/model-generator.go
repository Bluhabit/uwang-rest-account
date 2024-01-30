package common

import "gorm.io/gen"

func GenerateEntity() {
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: "./entity",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(GetDbConnection())
	g.GenerateAllTable()
	g.Execute()
}
