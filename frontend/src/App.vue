<template>
    <div>
        <ul id="example-1">
            <li v-for="item in this.people" :key="item.name">
                <person v-bind:person="item" />
            </li>
        </ul>
    </div>
</template>


<script>
import axios from 'axios';
import Person from './Person.vue';

export default {
  components: { Person },
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