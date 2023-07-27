<script lang="ts">
    import { onMount } from 'svelte';
    import { http } from "@astric";
	  import Tareas from './components/Tareas.svelte';
    import Textfield from '@smui/textfield';
    import CharacterCounter from '@smui/textfield/character-counter';
    import Fab, { Icon } from '@smui/fab';

     let tareas:any[] = [];

    let valueA = '';
    
    function crearProyecto(nombre:string) {
      let proyectos = {
        nombre: nombre,
        prioridad: 'MEDIA',
        fechaInicio: '29/06/23',
        fechaFin: '10/07/23'
      }
      http.post(`Proyectos/CrearProyecto`, proyectos)
        .then((data:any) => {
          console.log('Proyecto agregado');
        })
        .catch(e => {
          console.log((e));
        
        }) 
    }

     function obtenerProyectos() {
       http.get(`Proyectos/ObtenerProyectos`)
       .then((data:any) => {
           console.log(`Proyecto obtenido: ${data}`);
           tareas = data.datos;
         })
         .catch(e => {
          console.log((e));
        
         }) 
   }

   onMount(async () => {
    obtenerProyectos()
  });

</script>

<div class="contenedor1">
  <h3>Lista de Proyectos</h3>
</div>

<div class="flexy">
  <div class="margins">
    <div>
      <Textfield bind:value={valueA} label="Crear Proyecto" input$maxlength={150} />
        <Fab mini color="primary">
          <Icon on:click={() => crearProyecto(valueA)} class="material-icons">add</Icon>
        </Fab>
      <CharacterCounter class="izq" slot="helper">0 / 150</CharacterCounter>
      
    </div>
  </div>
  
  {#each tareas as item}
    <Tareas tarea={item} />    
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
</style>