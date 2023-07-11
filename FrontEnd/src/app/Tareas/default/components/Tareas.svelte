<script lang="ts">
	  import Separator from '@smui/list';
	  import Prioridad from './prioridad.svelte';
	  import Fecha from './fecha.svelte';
    import SubTarea from './subTarea.svelte';
    import Basura from './basura.svelte';
    import Paper from '@smui/paper';
    import Textfield from '@smui/paper';
    import Nota from "./nota.svelte";
    
    let value = '';
    export let tarea: any={};
</script>
  
<div >
    <Paper color="custom-red">
      <div class="titulo">
        <Textfield disabled bind:value={value}>{tarea.nombre}</Textfield>
        <Separator />
      </div>
      
      
      <div class="margins"> 
        <Nota></Nota>
        <SubTarea></SubTarea>
        <div class="columns margins forma">
          <Fecha>{tarea.fechaInicio}</Fecha>
          <Prioridad>{tarea.prioridad}</Prioridad>
          <Basura></Basura>
        </div>   
      </div>
    </Paper>
</div>

<style>
  :global(.custom-color) {
    border-color: #7A3E48;
    color: #5F021F !important;
    border-radius: 30px;
    border: 55px;
    margin: 2rem;
  }
  * :global(.shaped-outlined),
  * :global(.shaped-outlined .mdc-select__anchor) {
    border-radius: 28px;
  }
  * :global(.shaped-outlined .mdc-text-field__input) {
    padding-left: 32px;
    padding-right: 0;
  }
  *
    :global(.shaped-outlined
      .mdc-notched-outline
      .mdc-notched-outline__leading) {
    border-radius: 28px 0 0 28px;
    width: 28px;
  }
  *
    :global(.shaped-outlined
      .mdc-notched-outline
      .mdc-notched-outline__trailing) {
    border-radius: 0 28px 28px 0;
  }
  * :global(.shaped-outlined .mdc-notched-outline .mdc-notched-outline__notch) {
    max-width: calc(100% - 28px * 2);
  }
 
  .forma {
    display: flex;
    align-items: center;
    width: 14rem;
    margin: 15px;
  }
  .titulo {
    font-size: x-large;
    width: 200px;
    
  }
</style>