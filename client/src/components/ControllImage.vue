<template>
  <div class="controllImage">
    <!-- sendingEventを追加 -->
    <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions" v-on:vdropzone-sending="sendingEvent" v-on:vdropzone-removed-file="removeEvent"></vue-dropzone>
  </div>
</template>

<script>
import vue2Dropzone from 'vue2-dropzone'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import axios from 'axios'

export default {
  name: 'ControllImage',
  data: function () {
    return {
      dropzoneOptions: {
        url: 'http://localhost:8888/images',
        method: 'post',
        addRemoveLinks: 'true'
      }
    }
  },
  components: {
    // ケバブケース(<vue-dropzone>)でレンダリング
    vueDropzone: vue2Dropzone
  },
  mounted: function() {
    // mountした時に、一覧を取得する処理を追加する
    axios.get('http://localhost:8888/images').then(res => {
      res.data.forEach(res => {
        let filename = res.path.replace('http://localhost:8888/', '')
        let id = filename.replace('.png', '')
        let file = {size: res.size, name: filename, type: "image/png", upload: {uuid: id}}

        // コードからformに画像をセットする
        this.$refs.myVueDropzone.manuallyAddFile(file, res.path)
      })
    }).catch(err => {
      console.error(err)
    })
  },
  methods: {
    // sendingEvent UUIDをサーバに送信し判定してもらう
    sendingEvent(file, xhr, formData) {
      // アップロード時に「UUID」を付与
      formData.append('uuid', file.upload.uuid)
    },
    // removeEvent 削除イベント
    removeEvent(file) {
      axios.delete(`http://localhost:8888/images/${file.upload.uuid}`).then(res => {
        console.log(res.data)
      }).catch(err => {
        console.error(err)
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
