package models

//获取每个项目的memberlist
func GetMemberList(projects []*Project) {
	//对列表中的每个项目查询参与该项目的全部成员
	for _, proj := range projects {
		var memberList []*UserData
		DB.Table("k_user").
			Select([]string{"k_user.k_user_id", "k_user.user_name", "k_user.head_shot_url"}).
			Joins("left join k_user_in_project on k_user.k_user_id = k_user_in_project.user_id").
			Where("k_user_in_project.project_id = ?", proj.ProjectID).
			Scan(&memberList)
		proj.MemberList = memberList
	}
}

//获取当前用户参与的项目
func GetJoinProjects(userID string) []*Project {
	var joinedProjects []*Project
	//从k_user_in_project中取出当前用户id所参与的全部项目，保存到对应结构体中
	jp := DB.Table("k_user_in_project").Select("project_id").
		Where("user_id = ?", userID).QueryExpr()
	//获取查询到的project的全部信息
	DB.Where("project_id IN (?)", jp).Find(&joinedProjects)
	//获取每个项目的memberlist，并将其保存
	GetMemberList(joinedProjects)

	return joinedProjects
}

//获取当前用户管理的项目
func GetManageProjects(userID string) []*Project {
	var managedProjects []*Project
	//从k_user_in_project中取出当前用户id所管理的全部项目，保存到对应结构体中
	//管理项目要求role_id>=3，在查询中加入该条件进行筛选
	jp := DB.Table("k_user_in_project").Select("project_id").
		Where("user_id = ? AND role_id >= 3", userID).QueryExpr()
	//获取查询到的project的全部信息
	DB.Where("project_id IN (?)", jp).Find(&managedProjects)
	//获取每个项目的memberlist
	GetMemberList(managedProjects)

	return managedProjects
}