<script>
    import Section from './section.svelte';
    let open = null;
    export let skillsSections = [];
    const expandKeywordList = () => {
        open = 'keywordList';
    }

    const expandSentenceList = () => {
        open = 'sentenceList';
    }

    const addSkillSection = () => {
        skillsSections = [...skillsSections, {name: '', keywords: []}];
    }
</script>

<style>
    div.skill-buttons{
        display: flex;
        flex-flow: row nowrap;
    }
    div.skill-buttons>button{
        font-size: 1em;
    }
    div.inputs>div.keyword-inputs{
        display: flex;
        flex-flow: column nowrap;
    }
    div.inputs>div.keyword-inputs>input.skills-section{
        width:35%;
        min-width: 300px;
    }
    div.inputs>div.keyword-inputs>input.skills-keyword{
        width: 25%;
        min-width: 280px;
    }
</style>

<Section title='Skills'>
    {#if !open}
        <div class='skill-buttons'>
            <button on:click={expandKeywordList}>
                Keyword List
            </button>
        
            <button on:click={expandSentenceList}>
                Sentence List
            </button>
        </div>
    {/if}

    
    <div class='inputs'>
        {#if open == 'keywordList'}
            <div class='keyword-inputs'>
                {#each skillsSections as skillsSection}
                    <div class='title'>Section:</div>
                    <input class='skills-section' type='text' bind:value={skillsSection.name}/>
                    <div class='title'>Keywords:</div>
                    {#each skillsSection.keywords as keyword}
                        <input class='skills-keyword' type='text' bind:value={keyword}/>
                    {/each}
                    <button on:click={()=>{skillsSection.keywords = [...skillsSection.keywords, '']}}>
                        Add Keyword
                    </button>
                {/each}
            </div>
            <button on:click={addSkillSection}>
                Add Skills Section
            </button>
        {/if}
        {#if open == 'sentenceList'}
            <div class='keyword-inputs'>
                {#each skillsSections as skillsSection}
                    <div class='title'>Sentence:</div>
                    <input class='skills-section' type='text' bind:value={skillsSection.name}/>
                    <div class='title'>Keywords:</div>
                    {#each skillsSection.keywords as keyword}
                        <input class='skills-keyword' type='text' bind:value={keyword}/>
                    {/each}
                    <button on:click={()=>{skillsSection.keywords = [...skillsSection.keywords, '']}}>
                        Add Keyword
                    </button>
                {/each}
            </div>
            <button on:click={addSkillSection}>
                Add Skills Section
            </button>
        {/if}
        

    </div>

</Section>