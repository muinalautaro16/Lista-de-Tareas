import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import tsconfigPaths from 'vite-tsconfig-paths'
import { resolve } from "path";
import config from '../FrontEnd/config.json';
import { child_process } from 'vite-plugin-child-process'
import path from 'path'

const envDev = {
  'process.env.FALLO_CONEXION': JSON.stringify(config.mensajes.fallo_conexion),
  'process.env.FALLO_RUTA': JSON.stringify(config.mensajes.fallo_ruta),
  'process.env.TIMEUOT': JSON.stringify(config.mensajes.fallo_timeuot),
  'process.env.DEFAULT': JSON.stringify(config.mensajes.fallo_default),
  'process.env.URL': JSON.stringify(config.developer.api_url),
  'process.env.WS_URL': JSON.stringify(config.developer.ws_url),
  'process.env.WS_EDABLE': JSON.stringify(config.developer.websocket),
  'process.env.WS_TIME': JSON.stringify(config.developer.ws_time),
  'process.env.DEV': true,
  'process.env.LOGUIN': false,
  'process.env.USER': JSON.stringify(config.developer.user),
  'process.env.REQUEST_TIMEOUT': JSON.stringify(config.developer.request_timeuot_sec)
}

const envLogin = {
  'process.env.FALLO_CONEXION': JSON.stringify(config.mensajes.fallo_conexion),
  'process.env.FALLO_RUTA': JSON.stringify(config.mensajes.fallo_ruta),
  'process.env.TIMEUOT': JSON.stringify(config.mensajes.fallo_timeuot),
  'process.env.DEFAULT': JSON.stringify(config.mensajes.fallo_default),
  'process.env.URL': JSON.stringify(config.developer.api_url),
  'process.env.WS_URL': JSON.stringify(config.developer.ws_url),
  'process.env.WS_EDABLE': JSON.stringify(config.developer.websocket),
  'process.env.WS_TIME': JSON.stringify(config.developer.ws_time),
  'process.env.DEV': true,
  'process.env.LOGUIN': true,
  'process.env.USER': JSON.stringify(config.developer.user),
  'process.env.REQUEST_TIMEOUT': JSON.stringify(config.developer.request_timeuot_sec)
}

const envProd = {
  'process.env.FALLO_CONEXION': JSON.stringify(config.mensajes.fallo_conexion),
  'process.env.FALLO_RUTA': JSON.stringify(config.mensajes.fallo_ruta),
  'process.env.TIMEUOT': JSON.stringify(config.mensajes.fallo_timeuot),
  'process.env.DEFAULT': JSON.stringify(config.mensajes.fallo_default),
  'process.env.URL': JSON.stringify(config.production.api_url),
  'process.env.WS_URL': JSON.stringify(config.developer.ws_url),
  'process.env.WS_EDABLE': JSON.stringify(config.developer.websocket),
  'process.env.WS_TIME': JSON.stringify(config.developer.ws_time),
  'process.env.DEV': false,
  'process.env.LOGUIN': true,
  'process.env.REQUEST_TIMEOUT': JSON.stringify(config.production.request_timeuot_sec)
}

const alias = {
  "@astric": `${path.resolve(__dirname, '../FrontEnd/shared/imports/controlers.ts')}`,
  "@astric-validate": `${path.resolve(__dirname, '../FrontEnd/shared/imports/validate.ts')}`,
  "@astric-login": `${path.resolve(__dirname, '../FrontEnd/shared/imports/login.ts')}`,
  "@astric-env": `${path.resolve(__dirname, '../FrontEnd/config.json')}`,
  "@astric-store": `${path.resolve(__dirname, '../FrontEnd/src/base/store.ts')}`,
  "@": `${path.resolve(__dirname, '../FrontEnd/src/app')}/`
}

const backend = {
  name: "BACKEND",
  command: ["air"],
  watch: [/BackEnd/]
}

const backendLoguin = {
  name: "BACKEND",
  command: ["go", "run", "main.go", "--modo=BeckEndTKN"],
  watch: [/BackEnd/]
}

export default defineConfig(({ command, mode }) => {
  if (command === 'serve') {
    switch (mode) {
      case "front":
        return {
          plugins: [
            svelte(),
            tsconfigPaths(),
          ],
          define: envDev,
          resolve: {
            alias: alias
          }
        }
      case "full":
        return {
          plugins: [
            svelte(),
            tsconfigPaths(),
            child_process(backend)
          ],
          define: envDev,
          resolve: {
            alias: alias
          }
        }
      case "login":
        return {
          plugins: [
            svelte(),
            tsconfigPaths(),
            child_process(backendLoguin)
          ],
          define: envLogin,
          resolve: {
            alias: alias
          }
        }
      default:
        console.log("--mode es requerido (front, full, login)")
        break;
    }
  } else {
    return {
      plugins: [
        svelte(),
        tsconfigPaths()
      ],
      define: envProd,
      resolve: {
        alias: alias
      }
    }
  }
})

/*

export default defineConfig({
  plugins: [
    svelte(),
    tsconfigPaths()
  ],
  define: env,
  resolve: {
    alias: {
      $astric: resolve("FrontEnd/shared/imports/controlers.ts"),
      $astricValidate: resolve("FrontEnd/shared/imports/validate.ts"),
      @astric-env: resolve("FrontEnd/config.json"),
      @astric-store: resolve("FrontEnd/src/base/store.ts"),
      $: resolve("/FrontEnd/src/app"),
    }
  }
}
)*/
