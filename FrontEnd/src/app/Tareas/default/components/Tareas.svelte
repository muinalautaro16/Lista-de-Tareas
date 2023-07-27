<script lang="ts">
	  import Fecha from './fecha.svelte';
    import SubTarea from './subTarea.svelte';
    import Basura from './basura.svelte';
    import Paper, {Content as ContentPaper}  from '@smui/paper';
    import Textfield from '@smui/paper';
    import Nota from "./nota.svelte";

    let value = '';
    import Select, { Option } from '@smui/select';
    
    // * dsd
    export let tarea: any={};
    let valueColor = '';
    let prioridades =  [
      {nombre: 'ALTA', color: '#5c2b29'},
      {nombre: 'MEDIA', color: '#003555'},
      {nombre: 'BAJA', color: '#00440f'}
    ];

    function colorPrioridad(color: string) {
        const colorPick:any = prioridades.find(c => c.nombre===color)
        return colorPick.color
    }

    function handleChange(e:any) {
      const selectedOption = e.target.value;
      console.log('Opción seleccionada:', selectedOption);
    } 

</script>
  
<div style="display:flex; flex-direction:column">
    <Paper style="background-color: {colorPrioridad(tarea.prioridad)};">
      <ContentPaper style="display:flex; flex-direction:column">
        <div class="titulo">
        <Textfield bind:value={value} >{tarea.nombre}</Textfield>
      </div>

      <div class="margins"> 
        <Nota></Nota>
        <SubTarea id_proyecto={tarea.id}></SubTarea>
        
          <div class=" forma">
              <div class="forma">
                <Fecha>{tarea.fechaInicio}</Fecha>
                <Select on:click={handleChange} class="shaped-outlined" variant="outlined" bind:value={valueColor} label="Prioridad">
                  <Option Value="" />
                  {#each prioridades as prioridad}
                    <Option value={prioridad.nombre}>
                      {prioridad.nombre}
                      <svg fill={prioridad.color} viewBox="0 0 120 120" version="1.1" xmlns="http://www.w3.org/2000/svg">
                        <circle cx="60" cy="60" r="10"/>
                      </svg>
                    </Option>
                  {/each}
                </Select>
              </div>
            <div style="align-self: center"> <Basura id_proyecto={tarea.id}></Basura></div>
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