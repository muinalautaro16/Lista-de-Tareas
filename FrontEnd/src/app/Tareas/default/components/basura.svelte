<script lang='ts'>
	
    import IconButton from '@smui/icon-button';
    import { http } from "@astric";
    export let id_proyecto: any;

    function btnEliminar() {
        http.delete(`Proyectos/EliminarProyecto/${id_proyecto}`)
        .then( () => {
            console.log('Proyecto elimando')
        })
        .catch( () => {
            console.log('No hemos podido eliminar el proyecto')
        })
    }
    
</script>

<div>
    <IconButton style="font-size: 50px" class="material-icons" aria-label="Trash" on:click={btnEliminar}>delete</IconButton>
</div> 
