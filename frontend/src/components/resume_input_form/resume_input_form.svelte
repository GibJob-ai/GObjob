<script>
    let skills = []
    import Skills from './skills.svelte';
    import Experience from './experience.svelte';
    import Education from './education.svelte';
    let skillsKeywords;
</script>

<style>
    h1{
        text-decoration: underline;
    }

</style>

<h1>Resume Input Form</h1>

<Skills bind:skillsKeywords/>
<Experience title='Work Experience'/>
<Experience title='Projects'/>
<Education/>