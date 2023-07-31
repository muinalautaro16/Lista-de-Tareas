<script lang="ts">
	
	import { onMount } from 'svelte';
  import Textfield from '@smui/textfield';
  import Button, { Label } from '@smui/button';  
  import List, { Item, Meta } from '@smui/list';
  import Checkbox from '@smui/checkbox';
  import IconButton from '@smui/icon-button';
  import { http } from "@astric";
  let changeEvent: CustomEvent<{ changedIndices: number[] }> | null;
 

  export let id_proyecto: any;
  let valueA = '';
  let etapas: any[] = [];
  let selected = [""];

  $: {
    selected = etapas.filter((etapa) => etapa.estado).map((etapa) => etapa.nombre);
  }



  function obtenerEtapas(){
    http.get(`Proyectos/ObtenerEtapas/${id_proyecto}`)
    .then((data:any)=>{
      etapas = [...data.datos];
    }).catch(e => {
      console.log(e);
      etapas=[...[]]
    })
  }

  onMount(async ()=>{
    obtenerEtapas();
  })

  function finalizada(id: number) {
  for (let i = 0; i < etapas.length; i++) {
    if (etapas[i]?.id === id) {
      etapas[i].estado = etapas[i].estado; // Cambiar entre true y false
        console.log(etapas[i].estado)
        console.log(typeof(etapas[i].estado))
        setTimeout(() => {
          modificarEtapa(id, etapas[i].nombre, etapas[i].estado);
        }, 1000);
        
    }
  }
}

/*function timerModificacion(id: any, nombre: any, estado: boolean){
  modificarEtapa(id, nombre, estado);
}*/

function modificarEtapa(id: any, nombre: any, estado: any){
  let etapa={
    id: id,
    id_proyecto: id_proyecto,
    nombre: nombre,
    estado: !estado
  }
  http.put(`Proyectos/ModificarEtapa`, etapa)
  .then((data: any)=>{
    console.log("Etapa modificada")
    console.log(data)
  }).catch(e => {
    console.log(etapa)
    console.log(e)
  })
}

function crearEtapa(nombre: string){
  let etapa={
    id_proyecto: id_proyecto,
    nombre: nombre,
    estado: false
  }
  http.post('Proyectos/CrearEtapa', etapa)
  .then((data:any)=>{
    console.log(`Etapa agregada: ${data}`)
    obtenerEtapas()
  }).catch(e => {
    console.log(e)
  })
}

function eliminarEtapa(id: any){
  http.delete(`Proyectos/EliminarEtapa/${id}`)
  .then((data: any)=>{
    console.log("Etapa eliminada" + data)
    obtenerEtapas()
  }).catch(e => {
    console.log(e)
    obtenerEtapas()
  })
}
</script>


<class class="columns margins">
  <List
  class="demo-list"
  checkList
  on:SMUIList:selectionChange={(event) => (changeEvent = event)}
  style="display: flex;  flex-direction:column; align-items:flex-start"
  >
  {#each etapas as etapa}
  <div style="display:flex">
    <Item>
      <Meta>
        <Checkbox on:change={()=>modificarEtapa(etapa.id, etapa.nombre, etapa.estado)} bind:checked={etapa.estado} /> 
      </Meta>
      {#if etapa.estado}
        <Textfield  bind:value={etapa.nombre} style="text-decoration: line-through" on:change={()=>modificarEtapa(etapa.id, etapa.nombre, etapa.estado)}/>
      {:else}
        <Textfield  bind:value={etapa.nombre} on:change={()=>modificarEtapa(etapa.id, etapa.nombre, etapa.estado)}/>
      {/if}
    </Item>
    <IconButton class="material-icons" aria-label="Trash" on:click={() => eliminarEtapa(etapa.id)}>delete</IconButton>
  </div>  
  {/each}  
</List>

<div  style="display: flex;">
  <Button style="color:white; margin:2.5vh; margin-left:0vh;" on:click={()=>crearEtapa(valueA)}>
    <Label style="font-size:5vh;">+</Label>
  </Button>
  
  <Textfield bind:value={valueA} label="Crear Etapa"/>

</div>
</class>
