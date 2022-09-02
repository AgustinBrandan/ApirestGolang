package handler

import (
	"fmt"
    "time"
	"net/http"
	//"tutorial/models"

	"github.com/gin-gonic/gin"
)

type evento struct{
	Titulo string `json:"titulo"`
	Descripcionc string `json:"descripcionc"`
	Descripcionl string  `json:"descripcionl"`
	Fecha string `json: "fecha"`
	Hora string `json:"hora"`
	Organizador string `json:"organizador"`
	Lugar string `json: "lugar"`
	Estado string `json:"estado"`
}
// Registro de eventos
var eventos =[]evento{
	{Titulo:"La Konga",Descripcionc:"Konga",Descripcionl:"La Konga viene a cordoba", Fecha:"28-08-2022",Hora:"23:00",Organizador:"Victor Morales",Lugar:"Quality Espacio",Estado:"Publicado"},
	{Titulo:"Miranda",Descripcionc:"miranda",Descripcionl:"Miranda viene a cordoba", Fecha:"1-09-2022",Hora:"20:00",Organizador:"Victor Morales",Lugar:"Quality Espacio",Estado:"Publicado"},
	{Titulo:"Ulises Bueno",Descripcionc:"Ulises",Descripcionl:"Ulises viene a cordoba", Fecha:"22-10-2022",Hora:"21:00",Organizador:"Victor Morales",Lugar:"Quality Espacio",Estado:"Borrador"},
}


func GetAll(context *gin.Context){
	context.IndentedJSON(http.StatusOK,eventos)
}

func AddEvento(context *gin.Context){
	var newEvento evento
// Llamar a BindJSON para vincular el JSON recibido a nuevo evento
	if err := context.BindJSON(&newEvento); err != nil {
        return
    }
    // Agregar el nuevo evento
    eventos = append(eventos, newEvento)
    context.IndentedJSON(http.StatusCreated, newEvento)
}


// Elmina un evento
func DeleteEvento(c *gin.Context) {
    titulo := c.Param("titulo")
    index := -1
    for i := 0; i < len(eventos); i++ {
        if eventos[i].Titulo == titulo {
            index = 1
        }
    }
    if index == -1 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Evento not found",
        })
        return
    }
    eventos = append(eventos[:index], eventos[index+1:]...)
    c.JSON(http.StatusOK, gin.H{
        "message": "evento has been deleted",
    })
}

//Devolver 1 evento filtrado por titulo
func GetEvento(c *gin.Context) {
    titulo := c.Param("titulo")

    for _, a := range eventos {
        if a.Titulo == titulo {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Evento not found"})
}
//Devolver 1 evento filtrado por Fecha
func GetEventof(c *gin.Context) {
    fecha := c.Param("fecha")

    for _, b := range eventos {
        if b.Fecha == fecha {
            c.IndentedJSON(http.StatusOK, b)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Evento not found"})
}

//Devolver 1 evento filtrado por Estado
func GetEventoEstado(c *gin.Context) {
    estado := c.Param("estado")

    for _, b := range eventos {
        if b.Estado == estado {
            c.IndentedJSON(http.StatusOK, b)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Evento not found"})
}

//Funcion que devuelve los Publicados
func GetPublicados(c *gin.Context) {
    
 
    for i := 0; i < len(eventos); i++ {
        if eventos[i].Estado == "Publicado" {
            listaPublicados := eventos [i]
            c.IndentedJSON(http.StatusOK, listaPublicados)
        }
    }
}


// Actualizar un registro
func UpdateEvento(c *gin.Context) {
    titulo := c.Param("titulo")
    var newEvento evento
    if err := c.ShouldBindJSON(&newEvento); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    index := -1
    for i := 0; i < len(eventos); i++ {
        if eventos[i].Titulo == titulo {
            index = 1
        }
    }
    if index == -1 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "eventos not found",
        })
        return
    }
    eventos[index] = newEvento
   
    c.JSON(http.StatusOK, newEvento)
}

//Funcion para inscribirme  , No Funciona
func Getasistir(c *gin.Context) {
    //	fecha actual
	dt := time.Now()    
    // Formato MM-DD-YYYY
        dt.Format("01-02-2006")

    for i := 0; i < len(eventos); i++ {
        listaPublicados := eventos [i]
        myDate, err := time.Parse("2006-01-02 ", eventos[i].Fecha)
        if err != nil {
            panic(err)
        }
        if (myDate.After(dt)){
            fmt.Println(myDate)
            
          c.IndentedJSON(http.StatusOK, listaPublicados)
        }
        c.IndentedJSON(http.StatusOK, listaPublicados)
    }
    

}
