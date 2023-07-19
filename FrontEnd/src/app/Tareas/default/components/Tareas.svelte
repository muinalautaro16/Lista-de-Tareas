<script lang="ts">
	  
	  import Prioridad from './prioridad.svelte';
	  import Fecha from './fecha.svelte';
    import SubTarea from './subTarea.svelte';
    import Basura from './basura.svelte';
    import Paper, {Content as ContentPaper}  from '@smui/paper';
    import Textfield from '@smui/paper';
    import Nota from "./nota.svelte";

    let value = '';
    export let tarea: any={};

</script>
  
<div style="display:flex; flex-direction:column">
    <Paper >
      <ContentPaper style="display:flex; flex-direction:column">
        <div class="titulo">
        <Textfield bind:value={value} >{tarea.nombre}</Textfield>
      </div>

      <div class="margins"> 
        <Nota></Nota>
        <SubTarea></SubTarea>
        <div style='display:flex; justify-content:space-between; flex-wrap:wrap'>
            <div class=" forma">
              <div><Fecha>{tarea.fechaInicio}</Fecha>
              <Prioridad>{tarea.prioridad}</Prioridad></div>
              <div > <Basura></Basura></div>
            </div>
          
        </div>   
      </div>
      </ContentPaper>
      
    </Paper>
</div>

<style>
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
    flex-wrap: wrap;
    justify-content: space-between;
    margin: 15px;
  }
  .titulo {
    font-size: x-large;
    width: 200px;
    
  }
</style>