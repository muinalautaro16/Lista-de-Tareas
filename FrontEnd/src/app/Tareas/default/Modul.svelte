<script lang="ts">
	
	  import Tareas from './components/Tareas.svelte';
    import Textfield from '@smui/textfield';
    import Fab, { Icon } from '@smui/fab';
    let tareas = [{
      nombre: 'tarea1',
      prioridad: 'Alta',
      fechaInicio: '29/06/23',
      fechaFin: '10/07/23'
    },

    {
      nombre: 'tarea1',
      prioridad: 'Alta',
      fechaInicio: '29/06/23',
      fechaFin: '10/07/23'
    }]

    let valueA = '';
    
</script>

<div class="contenedor1">
  <h3>Lista de tareas</h3>
</div>

<div class="flexy"></div>
  <div class="margins">
    <div>
      <Textfield bind:value={valueA} label="Crear Tarea" />
      <Fab mini color="primary">
          <Icon class="material-icons">add</Icon>
      </Fab>
    </div>
  </div>
  {#each tareas as item}
    <Tareas tarea={item}>
    
    </Tareas>
  {/each}


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