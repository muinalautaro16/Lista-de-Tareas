<script lang="ts">
    import Textfield from '@smui/textfield';
    import CharacterCounter from '@smui/textfield/character-counter';
    import { createEventDispatcher } from 'svelte';
    export let valueNota: string;
    let dispatch = createEventDispatcher();
    $: nota = valueNota ? valueNota : '';

    function enviarNota() {
        dispatch('enviarNota', { nota: nota });
    }
</script>

<Textfield
    style="width: 100%;"
    textarea
    bind:value={nota}
    label="Nota"
    input$maxlength={500}
    on:change={() => {
        enviarNota();
    }}
>
    <CharacterCounter class="izq" slot="helper">0 / 500</CharacterCounter>
</Textfield>

<style>
</style>
