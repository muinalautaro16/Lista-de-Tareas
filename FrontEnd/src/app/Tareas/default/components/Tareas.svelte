<script lang="ts">
    import List, { Item, Separator, Text } from '@smui/list';
    import Checkbox from '@smui/checkbox';
    import FormField from '@smui/form-field';
    import Fab, { Icon } from '@smui/fab';
    import IconButton from '@smui/icon-button';
    import Paper, { Title, Content } from '@smui/paper';
    import Textfield from '@smui/textfield';
    import { DateInput } from 'date-picker-svelte';
    let date = new Date();
    let date2 = new Date();
    import Accordion, { Panel, Header } from '@smui-extra/accordion';
    import Select, { Option } from '@smui/select';
    let prioridades =  ['Alta', 'Media', 'Baja'];
    
    let value = '';
    let ValueB = '';
    let ValueC = ''
    let checked = false;

    export let tarea: any={};
</script>
  
<div >
    <Paper color="custom-red">
      <Title>{tarea.nombre}</Title>
      <div class="margins">
        <Textfield
          style="width: 100%;"
          textarea
          bind:value
          label="Nota"
        >
        </Textfield>
        <div style="color" class="accordion-container">
          <Accordion>
            <Panel>
              <Header>Subtareas</Header>
              <Content>
                <div class="subtarea">
                  <List class="demo-list" nonInteractive>
                    
                    <FormField>
                      <Checkbox bind:checked />
                      <Item><Text>Subtarea 1</Text></Item>
                    </FormField>
                    <Separator />
                    
                  </List>
                  <IconButton class="material-icons delete" aria-label="Trash">delete</IconButton>
                </div>
                
                <div>                 
                  <Fab mini style="color: black">
                      <Icon class="material-icons agregar">add</Icon>
                  </Fab>
                  <Textfield bind:value={ValueB} label="Crear Subtarea" />
                </div>
              </Content>
            </Panel>
          </Accordion>
        </div>  
        <div class="columns margins forma">
          <div>
            <Textfield
            type="date"
            class="shaped-outlined"
            variant="outlined"
            bind:value={date2}
            label="Label"
            >
              <Icon class="material-icons" slot="leadingIcon">event</Icon>
            </Textfield>
          </div>

          <div class="calendario forma">
            
            <Icon class="material-icons" slot="leadingIcon">event</Icon>
            <DateInput  bind:value={date} /> 
          </div>

          <div class="forma">
            <Select class="shaped-outlined" variant="outlined" bind:value={ValueC} label="Prioridad">
              <Option value="" />
              {#each prioridades as prioridad}
                <Option value={prioridad}>{prioridad}</Option>
              {/each}
            </Select>
          </div>
          <div class="delete" >
            <IconButton class="material-icons delete" aria-label="Trash">delete</IconButton>
          </div> 
        </div>
               
      </div>
    </Paper>
</div>

<style>
  :global(.custom-color) {
    border-color: #7A3E48;
    color: #5F021F !important;
    border-radius: 30px;
    border: 55px;
    margin: 2rem;
  }

  * :global(.shaped-outlined),
  * :global(.shaped-outlined .mdc-select__anchor) {
    border-radius: 28px;
  }
  * :global(.shaped-outlined .mdc-text-field__input) {
    padding-left: 32px;
    padding-right: 0;
  }
  *
    :global(.shaped-outlined
      .mdc-notched-outline
      .mdc-notched-outline__leading) {
    border-radius: 28px 0 0 28px;
    width: 28px;
  }
  *
    :global(.shaped-outlined
      .mdc-notched-outline
      .mdc-notched-outline__trailing) {
    border-radius: 0 28px 28px 0;
  }
  * :global(.shaped-outlined .mdc-notched-outline .mdc-notched-outline__notch) {
    max-width: calc(100% - 28px * 2);
  }
  *
    :global(.shaped-outlined.mdc-select--with-leading-icon
      .mdc-notched-outline:not(.mdc-notched-outline--notched)
      .mdc-floating-label) {
    left: 16px;
  }

  .calendario {
    display: contents;

  }

  .forma {
    display: flex;
    align-items: center;
    width: 14rem;
    margin: 15px;
  }

  * :global(.demo-list) {
    max-width: 240px;
    
  }

  .subtarea {
    display: flex;
    align-items: center;
  }

  .agregar {
    background-color: #1b1b1b;           
    
  }

  .delete {
    display: grid;
    justify-content: end;
    
  }
</style>