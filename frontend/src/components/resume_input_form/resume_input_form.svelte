<script>
    let skills = [];
    import Skills from './sections/skills/skills.svelte';
    import Experience from './sections/experience/experience.svelte';
    import Education from './sections/education/education.svelte';
    import Awards from './sections/awards/awards.svelte';
    let skillsKeywords;
</script>

<style src='./resume_input_form.scss'></style>

<h1>Resume Input Form</h1>

<Skills bind:skillsKeywords/>
<Experience title='Work Experience'/>
<Experience title='Projects'/>
<Education/>
<Awards/>