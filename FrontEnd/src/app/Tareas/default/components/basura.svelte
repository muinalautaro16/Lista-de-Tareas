<script>
	
    import IconButton from '@smui/icon-button';
    import { http } from "@astric";
    export let id_proyecto:any;

    function btnEliminar() {
        http.delete(`proyecto/EliminarProyecto/${id_proyecto}`)
        .then( () => {
            console.log('Proyecto elimando')
        })
        .catch( () => {
            console.log('No hemos podido eliminar el proyecto')
        })
    }
    // btnEliminar()
</script>

<div>
    <IconButton on:click{btnEliminar} style="font-size: 50px" class="material-icons" aria-label="Trash">delete</IconButton>
</div> 
