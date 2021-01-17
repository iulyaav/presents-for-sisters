<template>
    <div class="person-container" id="person-{{person.name}}">
        Hello, {{ person.name }}
    </div>
</template>

<script>
module.exports = {
  props: {'person': {
      type: Object
  }},
  data: function() {
    return {
      greeting: "Hello"
    };
  }
};
</script>
