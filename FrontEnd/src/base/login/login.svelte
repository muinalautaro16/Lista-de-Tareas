<script lang="ts">
    import FormField from '@smui/form-field';
    import Textfield from '@smui/textfield';
    import HelperText from '@smui/textfield/helper-text';
</script>

<FormField>
    <div>
        <Textfield variant="outlined" value="" label="Label">
            <HelperText slot="helper">Helper Text</HelperText>
        </Textfield>
    </div>
</FormField>
