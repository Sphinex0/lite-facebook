package repository

func (data *Database) SaveMember(member models.Member) (err error) { 
	args := utils.GetExecFields(member, "ID")
    _, err := data.Db.Exec(fmt.Sprintf(`
        INSERT INTO members
        VALUES (NULL, %v)
    `, utils.Placeholders(len(args))), args...)                                                                                           
    return
}
