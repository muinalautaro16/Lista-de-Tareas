<script lang="ts">
    import Sider from './Sider.svelte';
    import Menu from './Menu.svelte';

    let menu = [
        {
            name: 'CHART',
            route: '/prueba/chart',
        },
        {
            name: 'FORMS',
            route: '/prueba/forms',
        },
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
