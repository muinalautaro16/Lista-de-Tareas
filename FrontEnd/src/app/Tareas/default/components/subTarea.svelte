<script>
    import Accordion, { Panel, Header, Content } from '@smui-extra/accordion';
    import List, { Item, Separator, Text } from '@smui/list';
    import Checkbox from '@smui/checkbox';
    import Fab, { Icon } from '@smui/fab';
    import Textfield from '@smui/textfield';
    import FormField from '@smui/form-field';
    import IconButton from '@smui/icon-button';
    let ValueB = '';
    let checked = false;
</script>

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
            <IconButton class="material-icons" aria-label="Trash">delete</IconButton>
          </div>
          
          <div>                 
            <Fab mini color="secondary">
                <Icon class="material-icons agregar">add</Icon>
            </Fab>
            <Textfield bind:value={ValueB} label="Crear Subtarea" />
          </div>
        </Content>
      </Panel>
    </Accordion>
  </div>  

<style>
.subtarea {
    display: flex;
    align-items: center;
}

* :global(.demo-list) {
    max-width: 240px;
    
  }
</style>