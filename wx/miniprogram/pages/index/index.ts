// index.ts
// 获取应用实例
const app = getApp<IAppOption>()

Page({
  isPageShowing:false,
  data: {
    // 默认值
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: '',
      layerStyle: -1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: true,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    location: {
      latitude: 23.099994,
      longitude: 113.324520,
    },
    scale: 10,
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: 23.099994,
        longitude: 113.324520,
        width: 50,
        height: 50
      }, {
        iconPath: "/resources/car.png",
        id: 1,
        latitude: 23.099994,
        longitude: 113.324520,
        width: 50,
        height: 50,
      },
    ],
    avatarURL:'',
  },
  onMyLocationTap() {
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        console.log(res)
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
          scale:20,
        })
      },
      fail: (res) => {
        console.log(res)
        wx.showToast({
          title: '请前往设置页授权',
          icon: 'none',
        })
      }
    })
  },
  onShow(){
    this.isPageShowing = true
  },
  async onLoad() {
    const userInfo = await getApp<IAppOption>().globalData.userInfo
    this.setData({
        avatarURL: userInfo.avatarUrl,
    })
  },
  onHide(){
    this.isPageShowing = false
  },
  moveCars(){
    const map = wx.createMapContext('map')
    const dest = {
      latitude:this.data.markers[0].latitude,
      longitude:this.data.markers[0].longitude,
    }

    const moveCar= ()=>{
      dest.latitude += 0.01
      dest.longitude += 0.01
      console.log(this.data.markers[0].latitude)
      map.translateMarker({
        destination:{
          latitude:dest.latitude,
          longitude:dest.longitude,
        },
        markerId:0,
        autoRotate:false,
        rotate:0,
        duration:5000,
        animationEnd:()=>{
          if(this.isPageShowing){
            moveCar()
          }
        },
      })
    }
    moveCar();
  },
  onScanClicked(){
    wx.scanCode({
      success:()=>{
        wx.navigateTo({
          url:'/pages/register/register',
        })
      },
      fail:console.error,
    })
  }
})
