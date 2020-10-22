
package main
import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/patiwatpanda/app/controllers"
	"github.com/patiwatpanda/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Employees struct {
	Employee []Employee
}

type Employee struct {
	EmployeeName     string
	EmployeeEmail    string
	EmployeePassword string
}

type Dentists struct {
	Dentist []Dentist
}

type Dentist struct {
	DentistName string
}
type Patients struct {
	Patient []Patient
}
type Patient struct {
	PatientName string
	PatientAge  int
}
type PillorderItems struct {
	PillorderItem []PillorderItem
}

type PillorderItem struct {
	PillorderItemName string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewPillorderController(v1, client)
	controllers.NewDentistderController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewPillorderItemController(v1, client)

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"patiwat fangsongnern", "chanwit@gmail.com", "123456"},
			Employee{"Name Surname", "Name@example.com", "123456"},
			Employee{"pratud chund", "pratud@gmail.com", "123456"},
			Employee{"way Roony", "way@example.com", "123456"},
		},
	}

	for _, e := range employees.Employee {
		client.Employee.
			Create().
			SetEmployeeName(e.EmployeeName).
			SetEmployeeEmail(e.EmployeeEmail).
			SetEmployeePassword(e.EmployeePassword).
			Save(context.Background())
	}

	// Set Dentist Data
	dentists := Dentists{
		Dentist: []Dentist{
			Dentist{"นพ.สมบัติ"},
			Dentist{"นพ.สมศรี"},
			Dentist{"นพ.สมสมัย"},
			Dentist{"นพ.สมหมาย"},
		},
	}
	for _, d := range dentists.Dentist {
		client.Dentist.
			Create().
			SetDentistName(d.DentistName).
			Save(context.Background())

	}
	//Set Patient Data
	patients := Patients{
		Patient: []Patient{
			{"Ben ya", 50},
			{"Somchai naka", 25},
			{"Somying boonmee", 50},
			{"Peenoy banmai", 29},
			{"Peelo poo", 24},
		},
	}
	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.PatientName).
			SetPatientAge(p.PatientAge).
			Save(context.Background())

	}

	//Set PillorderItem Data
	pillorderitems := PillorderItems{
		PillorderItem: []PillorderItem{
			PillorderItem{"แก้ไข้"},
			PillorderItem{"แก้ปวด"},
			PillorderItem{"พาราเซลตามอล"},
			PillorderItem{"คลายกล้ามเนื้อ"},
			PillorderItem{"แก้ไอ"},
			PillorderItem{"ชูกำลัง"},
			PillorderItem{"น้ำ"},
		},
	}
	for _, pi := range pillorderitems.PillorderItem {
		client.PillorderItem.
			Create().
			SetPillorderItemName(pi.PillorderItemName).
			Save(context.Background())

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}