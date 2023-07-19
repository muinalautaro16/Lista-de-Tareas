<script>
    import Textfield from '@smui/textfield';
    import CharacterCounter from '@smui/textfield/character-counter';
    let value = '';
</script>

<Textfield style="width: 100%;" textarea bind:value label="Nota" input$maxlength={500}>
  <CharacterCounter class="izq" slot="helper">0 / 500</CharacterCounter>
</Textfield>

<style>
 
</style>