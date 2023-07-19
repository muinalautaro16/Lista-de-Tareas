<script lang="ts">
	import ListaSubtarea from './listaSubtarea.svelte';
  import Accordion, { Panel, Header, Content } from '@smui-extra/accordion';
  import IconButton, { Icon } from '@smui/icon-button';
  
  let contadorTareas = 2;
  let contadorTareasListas= 0;

  let panel1Open = false;
</script>

<div class="accordion-container">
  <Accordion>
    <Panel bind:open={panel1Open}>
      <Header style="display:flex-direcction; flex-direction:row-reverse; background-color:#272822">
        <IconButton slot="icon" toggle pressed={panel1Open}>
          <Icon class="material-icons" on>expand_less</Icon>
          <Icon class="material-icons">expand_more</Icon>
        </IconButton>
        Etapas - {contadorTareasListas} de {contadorTareas} Terminadas
      </Header>
      <Content>
        <ListaSubtarea/>
      </Content>
    </Panel>
  </Accordion>
</div>
