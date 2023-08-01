<script lang="ts">
    import { onMount } from 'svelte';
    import { http } from '@astric';
    import Tareas from './components/Tareas.svelte';
    import Textfield from '@smui/textfield';
    import CharacterCounter from '@smui/textfield/character-counter';
    import Fab, { Icon } from '@smui/fab';

    let tareas: any[] = [];
    let proyecto: any;
    let today = new Date();
    let todayString =
        today.getFullYear() +
        '-' +
        ('0' + (today.getMonth() + 1)).slice(-2) +
        '-' +
        ('0' + today.getDate()).slice(-2) +
        'T' +
        +('0' + today.getHours()).slice(-2) +
        ':' +
        ('0' + today.getMinutes()).slice(-2);

    let valueA = '';

    function proyectoRecibido(event: any) {
        proyecto = event.detail.datos;
        modificarProyecto();
    }

    function modificarProyecto() {
        http.put('Proyectos/ModificarProyecto', proyecto)
            .then((data: any) => {
                console.log('proyecto modificado');
                console.log(data);
            })
            .catch(e => {
                console.log(e);
            });
    }

    function crearProyecto(nombre: string) {
        let proyectos = {
            nombre: nombre,
            nota: '',
            prioridad: 'MEDIA',
            fecha: todayString,
        };
        http.post(`Proyectos/CrearProyecto`, proyectos)
            .then((data: any) => {
                console.log('Proyecto agregado');
                obtenerProyectos();
                console.log(data);
            })
            .catch(e => {
                console.log(e);
            });
    }

    function obtenerProyectos() {
        http.get(`Proyectos/ObtenerProyectos`)
            .then((data: any) => {
                console.log(`Proyecto obtenido: ${data}`);
                tareas = data.datos;
            })
            .catch(e => {
                console.log(e);
            });
    }

    onMount(async () => {
        obtenerProyectos();
    });
</script>

<div class="contenedor1">
    <h3>Lista de Proyectos</h3>
</div>

<div class="flexy">
    <div class="margins">
        <div class="textfield">
            <div>
                <Textfield
                    bind:value={valueA}
                    label="Crear Proyecto"
                    input$maxlength={150}
                >
                    <CharacterCounter slot="helper">0 / 150</CharacterCounter>
                </Textfield>
            </div>

            <div class="boton">
                <Fab mini color="primary">
                    <Icon
                        on:click={() => crearProyecto(valueA)}
                        class="material-icons">add</Icon
                    >
                </Fab>
            </div>
        </div>
    </div>

    {#each tareas as item}
        <Tareas
            on:proyectoEnviado={proyectoRecibido}
            on:eliminado={proyEl => {
                tareas = tareas.filter(proyArray => {
                    proyArray.id !== proyEl.detail.delete;
                });
                obtenerProyectos();
            }}
            tarea={item}
        />
    {/each}
</div>

<style>
    * :global(svg:focus) {
        outline: 0;
    }

    .contenedor1 {
        display: grid;
        text-align: center;
        justify-content: center;
        width: 100%;
        background-color: #3c3c3c;
        margin: 10px;
        height: auto;
    }

    .textfield {
        display: flex;
        justify-items: center;
    }

    .boton {
        display: grid;
        align-items: center;
        padding-left: 8px;
        padding-bottom: 5px;
    }
</style>
