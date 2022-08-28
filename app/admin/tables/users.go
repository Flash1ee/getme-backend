package tables

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserTable(ctx *context.Context) (userTable table.Table) {

	// config the table model.
	userTable = table.NewDefaultTable(table.Config{
		Driver:     db.DriverPostgresql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Bigint,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := userTable.GetInfo()

	// set id sortable.
	info.AddField("ID", "id", db.Bigint).FieldSortable()
	info.AddField("FirstName", "first_name", db.Text)
	info.AddField("LastName", "last_name", db.Text)
	info.AddField("Nickname", "nickname", db.Text)
	info.AddField("About", "about", db.Text)
	info.AddField("Avatar", "avatar", db.Text)
	info.AddField("IsSearchable", "is_searchable", db.Boolean)
	info.AddField("CreatedAt", "created_at", db.Timestamptz)
	info.AddField("UpdatedAt", "updated_at", db.Timestamptz)
	info.AddField("TelegramTag", "tg_tag", db.Text)

	// set the title and description of table page.
	info.SetTable("users").SetTitle("Users").SetDescription("Application users")
	//SetAction(template.HTML(`<a href="http://google.com"><i class="fa fa-google"></i></a>`)) // custom operation button

	formList := userTable.GetForm()

	// set id editable is false.
	formList.AddField("ID", "id", db.Int, form.Default)
	formList.AddField("FirstName", "first_name", db.Text, form.Text)
	formList.AddField("LastName", "last_name", db.Text, form.Text)
	formList.AddField("Nickname", "nickname", db.Text, form.Text)
	formList.AddField("About", "about", db.Text, form.Text)
	formList.AddField("Avatar", "avatar", db.Text, form.File)
	formList.AddField("IsSearchable", "is_searchable", db.Boolean, form.Default).FieldDefault("false")
	formList.AddField("TelegramTag", "tg_tag", db.Text, form.Text)

	// set the title and description of form page.
	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	// use SetPostHook to add operation when form posted.
	formList.SetPostHook(func(values form2.Values) error {
		fmt.Println("userTable.GetForm().PostHook", values)
		return nil
	})

	return
}
