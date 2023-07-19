<script>
	
    import Select, { Option } from '@smui/select';
    let valueC = ''
    let prioridades =  [
      {nombre: 'ALTA', color: '#5c2b29'},
      {nombre: 'MEDIA', color: '#003555'},
      {nombre: 'BAJA', color: '#00440f'}
    ];

</script>

<div>
    <Select class="shaped-outlined" variant="outlined" bind:value={valueC} label="Prioridad">
      <Option Value="" />
      {#each prioridades as prioridad}
        <Option value={prioridad}>
          {prioridad.nombre}
          <svg fill={prioridad.color} viewBox="0 0 120 120" version="1.1" xmlns="http://www.w3.org/2000/svg">
            <circle cx="60" cy="60" r="10"/>
          </svg>
        </Option>
      {/each}
    </Select>
</div>

<style>

</style>