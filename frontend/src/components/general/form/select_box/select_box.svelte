<script>
	export let items;
	// ACCEPTED ITEMS FORMATS:
	// -----FLAT:
	// [{'id':'item1', 'title': 'Item 1'}, ...]
	// -----GROUPED:
	// [{'id':'group1', 'title': 'Group 1', 'items': [{'id':'item1', 'title': 'Item 1'}]}, ...]
	export let label; // ie. 'select resume'
	export let selected; // selected item id (binding)
	export let expanded = false;

	let active=false;

	$: grouped = items[0].hasOwnProperty(items);

	const updateSelection = (id) => {
		selected = id;
	}
	const selectedTitle = () => {
		items.forEach(function (i) {
				if (i.id === selected){
						return i.title;
				}
		});
		return label;
	}
// $('select').each(function(){
    // var $this = $(this), numberOfOptions = $(this).children('option').length;
  
    // $this.addClass('select-hidden'); 
    // $this.wrap('<div class="select"></div>');
    // $this.after('<div class="select-styled"></div>');

    // var $styledSelect = $this.next('div.select-styled');
    // $styledSelect.text($this.children('option').eq(0).text());
  
    // var $list = $('<ul />', {
        // 'class': 'select-options'
    // }).insertAfter($styledSelect);
  
    // for (var i = 0; i < numberOfOptions; i++) {
        // $('<li />', {
            // text: $this.children('option').eq(i).text(),
            // rel: $this.children('option').eq(i).val()
        // }).appendTo($list);
    // }
  
    // var $listItems = $list.children('li');
  
    // $styledSelect.click(function(e) {
        // e.stopPropagation();
        // $('div.select-styled.active').not(this).each(function(){
            // $(this).removeClass('active').next('ul.select-options').hide();
        // });
        // $(this).toggleClass('active').next('ul.select-options').toggle();
    // });
  
    // $listItems.click(function(e) {
        // e.stopPropagation();
        // $styledSelect.text($(this).text()).removeClass('active');
        // $this.val($(this).attr('rel'));
        // $list.hide();
        // //console.log($this.val());
    // });
  
    // $(document).click(function() {
        // $styledSelect.removeClass('active');
        // $list.hide();
    // });
// });
</script>
<style src='select_box2.scss'>
</style>
<div class='select'>
	<select bind:value={selected}
					size={expanded?6:null}
					class:expanded
					class="select-hidden">
		<option value="" disabled selected hidden>
			{label}
		</option>
		{#if grouped}
			{#each items as group}
				<optgroup label="{group.title}">
					{#each group.items as item}
						<option class='item indented' value={item.id}>
							{item.title}
						</option>
					{/each}
				</optgroup>
			{/each}
		{:else}
			{#each items as item}
				<option class='item' value={item.id}>
					{item.title}
				</option>
			{/each}
		{/if}
	</select>
	<!-- <span class="focus"></span> -->
	<div class='select-styled'
			 on:click={()=>{active=!active}}
			 class:active>
		{selectedTitle()}
	</div>
	{#if active}
		<ul class='select-options'>
			{#each items as item}
				<li on:click={()=>{active=false}}>
					{item.title}
				</li>
			{/each}
		</ul>
	{/if}
</div>
