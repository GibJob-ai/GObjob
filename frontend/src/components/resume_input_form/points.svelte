<script>
    export let points;

    const addPoint = () => {
        points = [...points, {point: '', keywords: []}];
    }

    const addKeyword = (point) => {
        point.keywords = [...point.keywords, ''];
    } 
</script>

<div class='points'>
    {#each points as point}
        <input type='text' bind:value={point.text}/>
        {#each point.keyword as keyword}
            <input type='text' bind:value={keyword}/>
        {/each}
        <button on:click={()=>{addKeyword(point)}}>
            Add Keyword
        </button>   
    {/each}
    <button on:click={addPoint}>
        Add Point
    </button>
</div>