<script>
    import Select, { Option } from '@smui/select';
    let ValueC = ''
    let prioridades =  ['Alta', 'Media', 'Baja'];
</script>

<div>
    <Select class="shaped-outlined" variant="outlined" bind:value={ValueC} label="Prioridad">
      <Option Value="" />
      {#each prioridades as prioridad}
        <Option value={prioridad}>{prioridad}</Option>
      {/each}
    </Select>
</div>

<style>

</style>