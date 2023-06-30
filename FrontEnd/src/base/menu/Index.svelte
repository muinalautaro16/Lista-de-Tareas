<script lang="ts">
    import Sider from './Sider.svelte';
    import Menu from './Menu.svelte';

    let menu = [
        {
            name: 'Proyecto',
            route: '/Tareas/default',
        }
    ];

    let databses = [
        {
            name: 'Suc1',
            db: 'db1',
        },
        {
            name: 'Suc2',
            db: 'db2',
        },
    ];
</script>

<Sider {databses} />
<Menu {menu} />
