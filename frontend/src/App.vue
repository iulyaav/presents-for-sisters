<template>
    <div>
        <ul id="example-1">
            <li v-for="item in this.people" :key="item.name">
                {{ item.name }}
            </li>
        </ul>
    </div>
</template>


<script>
import axios from 'axios';

export default {
  name: 'App',

  data() { return {
    people: '',
  } },
  created() {
    axios.get('http://localhost:3000/').then(people => {
      this.people = people.data.persons;
    });
  },

  methods: {
    getData() {
      console.log(`I should create a website thumbnail of ${this.websiteUrl}`);
    }
  }
}
</script>