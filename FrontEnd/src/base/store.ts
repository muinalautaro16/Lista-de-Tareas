import { writable } from 'svelte/store';
import { env } from '../../ASTRIC/config/env'
import { db } from '@astric-env'


let menu = {
    open: false,
}

let socket = {
    //App sirve como identificador de app
    app: "",
    //EndPoint sirve como identificador del endpoint
    endPoint: "",
    //Tag accion o transaccion realizada
    tag: "",
    //BD base de datos en la cual se produjo el cambio
    BD: "",
    //Tabla de la base de datos en la cual se produjo el cambio
    tabla: "",
    //Usuario usuario que produjo el evento
    usuario: "",
    //Data datos enviados si es nesesario
    data: "",
    //Msj mensaje si es nesesario
    msj: "",
    //Error si es nesessario
    error: ""
}

export let dbStore = {
    db: db.default,
    name: db.name,
}

let login = {
    loguin: !env.login
}

export let wsAccion = {
    set: (dato: any) => {
        ws.update(value => {
            const datos = JSON.parse(dato.data)
            value.app = datos.App
            value.BD = datos.BD
            value.data = datos.Data
            value.endPoint = datos.EndPoint
            value.error = datos.Error
            value.msj = datos.Msj
            value.tabla = datos.Tabla
            value.tag = datos.Tag
            value.usuario = datos.Usuario
            return value
        })
    }
}

export let menuAccion = {
    open: () => {
        menuStore.update(value => {
            value.open = !value.open
            return value
        })
    }
}

export let dbAccion = {
    set: (name: string, db: string) => {
        dbStore.db = db
        database.update(value => {
            value.db = db
            value.name = name
            return value
        })
    }
}

export let loguinAccion = {
    login: () => {
        loguin.update(value => {
            value.loguin = true
            return value
        })
    },
    unLoguin: () => {
        loguin.update(value => {
            value.loguin = false
            return value
        })
    }
}

export let ws = writable(socket)
export let loguin = writable(login)
export let database = writable(dbStore)
export let menuStore = writable(menu)
