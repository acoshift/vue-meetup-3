<template>
<div>
	<ul v-if="list">
		<li v-for="x in list" :key="x">{{x}}</li>
	</ul>
</div>
</template>

<script>
export default {
	props: {
		list: Array
	}
}
</script>
