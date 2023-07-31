<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import IconButton from '@smui/icon-button';
    import { http } from '@astric';
    export let id_proyecto: any;

    const dispatch = createEventDispatcher();

    function btnEliminar() {
        http.delete(`Proyectos/EliminarProyecto/${id_proyecto}`)
            .then(() => {
                console.log('Proyecto elimando');
                dispatch('eliminado', {
                    delete: id_proyecto,
                });
            })
            .catch(() => {
                console.log('No hemos podido eliminar el proyecto');
            });
    }
</script>

<div>
    <IconButton
        style="font-size: 50px"
        class="material-icons"
        aria-label="Trash"
        on:click={btnEliminar}>delete</IconButton
    >
</div>
