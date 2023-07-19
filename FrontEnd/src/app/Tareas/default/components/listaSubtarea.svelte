<script lang="ts">
  import Textfield from '@smui/textfield';
  import Button, { Label } from '@smui/button';  
  import List, { Item, Meta } from '@smui/list';
  import Checkbox from '@smui/checkbox';
  import IconButton from '@smui/icon-button';

  let changeEvent: CustomEvent<{ changedIndices: number[] }> | null;

  let valueA = '';

  let selected = [""];
</script>


<class class="columns margins">
  <List
  class="demo-list"
  checkList
  on:SMUIList:selectionChange={(event) => (changeEvent = event)}
  style="display: flex;  flex-direction:column; align-items:flex-start"
  >
  <div style="display:flex">
    <Item>
      <Meta>
        <Checkbox bind:group={selected} value="Etapa 1" />
      </Meta>
      <Label>Etapa 1</Label>      
    </Item>
    <IconButton class="material-icons" aria-label="Trash">delete</IconButton>
  </div>
  <div style="display:flex">
    <Item>
      <Meta>
        <Checkbox bind:group={selected} value="Etapa 2" />
      </Meta>
      <Label>Etapa 1</Label>      
    </Item>
    <IconButton class="material-icons" aria-label="Trash">delete</IconButton>
  </div>
</List>

<div  style="display: flex;">
  <Button style="color:white; margin:2.5vh; margin-left:0vh;">
    <Label style="font-size:5vh;">+</Label>
  </Button>
  
  <Textfield bind:value={valueA} label="Crear Etapa"/>    
</div>
</class>
