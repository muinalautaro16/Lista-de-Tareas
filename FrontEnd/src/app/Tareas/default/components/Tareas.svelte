<script lang="ts">
	  import Fecha from './fecha.svelte';
    import SubTarea from './subTarea.svelte';
    import Basura from './basura.svelte';
    import Paper, {Content as ContentPaper, Title}  from '@smui/paper';
    import Textfield from '@smui/textfield';
    import Nota from "./nota.svelte";
    import Select, { Option } from '@smui/select';
    
    // * dsd
    export let tarea: any={};

    let prioridades =  [
      {nombre: 'ALTA', color: '#5c2b29', borde: '#281312'},
      {nombre: 'MEDIA', color: '#003555', borde: '#001521'},
      {nombre: 'BAJA', color: '#00440f', borde: '#003606'}
    ];

    function colorPrioridad(color: string) {
        const colorPick:any = prioridades.find(c => c.nombre===color)
        return colorPick
    }

    function handleChange(e:any) {
      const selectedOption = e.target.value;
      console.log('Opción seleccionada:', selectedOption);
    } 

    function modificarProyecto(data:any) {
      http.put("Proyectos/ModificarProyecto")
      .then((data:any) => {
        console.log("Proyecto modificado");
      })
    }

    
</script>
  
<div style="display:flex; flex-direction:column; margin: 1rem;">
    <Paper style="--borde:{colorPrioridad(tarea.prioridad).borde}; background-color: {colorPrioridad(tarea.prioridad).color}; border-radius: 5px 45px; border:10px solid var(--borde)">
      <Title style="padding-bottom: 10px">
          <Textfield bind:value={tarea.nombre}></Textfield>
      </Title>

      <ContentPaper style="display:flex; flex-direction:column">

        <div class="margins"> 
          <Nota></Nota>
          <SubTarea id_proyecto={tarea.id}></SubTarea>
          
              <div class=" forma">
                <div class="forma">
                  <Fecha>{tarea.fechaInicio}</Fecha>
                  <Select class="shaped-outlined" variant="outlined" bind:value={tarea.prioridad} label="Prioridad">
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
              <div style="align-self: center"> <Basura id_proyecto={tarea.id} ></Basura></div>
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
</style>