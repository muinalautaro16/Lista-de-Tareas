<script lang="ts">
    import Paper, { Title, Content } from '@smui/paper';
    import Textfield from '@smui/textfield';
    import Icon from '@smui/textfield/icon';
    import { DateInput } from 'date-picker-svelte';
    let date = new Date();
    import Accordion, { Panel, Header } from '@smui-extra/accordion';
    import Select, { Option } from '@smui/select';
    let prioridades =  ['Alta', 'Media', 'Baja'];
    
    let value = '';

    export let tarea: any={};
</script>
  
<div class="paper-container">
    <Paper color="custom-green">
      <Title>{tarea.nombre}</Title>
      <div class="margins">
        <Textfield
          style="width: 100%;"
          textarea
          bind:value
          label="Nota"
        >
        </Textfield>
        <div class="accordion-container">
          <Accordion>
            <Panel>
              <Header>Subtareas</Header>
              <Content>
                The content for panel 1.
         
                <ul>
                  <li>Some</li>
                  <li>List</li>
                  <li>Items</li>
                </ul>
              </Content>
            </Panel>
          </Accordion>
        </div>  
        <div class="columns margins forma">
          <div class="calendario">
            <Icon class="material-icons" slot="leadingIcon">event</Icon>
            <DateInput class="shaped-outlined" bind:value={date} />
          </div>

          <div class="forma">
            <Select class="shaped-outlined" variant="outlined" bind:value label="Prioridad">
              <Option value="" />
              {#each prioridades as prioridad}
                <Option value={prioridad}>{prioridad}</Option>
              {/each}
            </Select>
          </div>
        </div>
        
      </div>
    </Paper>
</div>

<style>
  
  .calendario {
    display: contents;

  }

  .forma {
    display: flex;
    align-items: center;
    width: 14rem;
    margin: 15px;
   
  }
</style>