<template>
  <div class="hello">
    <!-- sendingEventを追加 -->
    <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions" v-on:vdropzone-sending="sendingEvent" v-on:vdropzone-removed-file="removeEvent"></vue-dropzone>
  </div>
</template>

<script>
import vue2Dropzone from 'vue2-dropzone'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import axios from "axios"

export default {
  name: 'HelloWorld',
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
    vueDropzone: vue2Dropzone
  },
  mounted: function() {
    axios.get('http://localhost:8888/images').then(res => {
      res.data.forEach(res => {
        // filename 取得
        let filename = res.path.replace('http://localhost:8888/', '')
        // uuid 取得
        let id = filename.replace('.png', '')
        // file オブジェクト作成
        var file = {size: res.size, name: filename, type: "image/png", upload: {uuid: id}}
        // コードからformに画像をセットする
        this.$refs.myVueDropzone.manuallyAddFile(file, res.path)
      })
    }).catch(err => {
      // 例外処理
      console.error(err)
    })
  },
  methods: {
    // UUIDをサーバに送信し判定してもらう
    sendingEvent: function (file, xhr, formData) {
      // アップ時に「UUID」を付与
      formData.append('uuid', file.upload.uuid)
    },
    // これで合っているのか謎
    removeEvent: function (file) {
    // removeEvent: function (file, error, xhr) {
      axios.delete(`http://localhost:8888/images/${file.upload.uuid}`).then(res => {
        console.log(res.data)
      }).catch(err => {
        console.error(err)
      })
    }
  }
}

// export default {
//   name: 'HelloWorld',
//   props: {
//     msg: String
//   }
// }
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
