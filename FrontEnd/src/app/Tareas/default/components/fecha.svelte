<script lang="ts">
    import Textfield from '@smui/textfield';
    import { createEventDispatcher } from 'svelte';
    export let valueFecha: string;
    $: fecha = valueFecha ? valueFecha : '';
    let dispatch = createEventDispatcher();

    function enviarFecha() {
        dispatch('fechaEnviada', { fecha: fecha });
    }
</script>

<div>
    <Textfield
        type="datetime-local"
        class="shaped-outlined"
        variant="outlined"
        bind:value={fecha}
        label="Fecha"
        on:change={() => {
            enviarFecha();
        }}
    />
</div>
