<script>
    import Icon from "@iconify/svelte";
    import { page } from '$app/stores';

    $: currentPath = $page.url.pathname;

    let menuItems = [
        { path: '/docs/introduction', icon: 'charm:home', label: 'Introduction' },
        { path: '/docs/screenshot', icon: 'fluent:screenshot-16-regular', label: 'Screenshot' },
    ];
</script>

<ul class='w-full flex flex-col space-y-2'>
    {#each menuItems as item}
        <li class={currentPath === item.path ? 'bg-[#e4e9ec] w-full rounded py-2 px-3' : 'w-full rounded py-2 px-3 hover:bg-[#e4e9ec] duration-150'}>
            <a href={item.path} class='flex items-center text-[#494953]'>
                <Icon icon={item.icon} class='mr-2' />
                <span>{item.label}</span>
            </a>
        </li>
    {/each}
</ul>