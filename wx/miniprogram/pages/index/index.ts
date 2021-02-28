// index.ts
// 获取应用实例
const app = getApp<IAppOption>()

Page({
  data: {
    motto: 'Hello TypeScript',
    userInfo: {},
    hasUserInfo: false,
    // canIUse: wx.canIUse('button.open-type.getUserInfo'),
  },
  // 事件处理函数
  bindViewTap() {
    wx.redirectTo({
      url: '../logs/logs',
    })
  },
  onLoad() {
app.globalData.userInfo.then(
      userInfo=>{
        this.setData({
          userInfo:userInfo,
          hasUserInfo: true,
        })
      }
    )
    //this.updateMotto()
  },
  getUserInfo(e: any) {
    console.log(e)
    const userInfo:WechatMiniprogram.UserInfo = e.detail.userInfo
    app.resoveUserInfo(userInfo)
  },

  //闭包应用
  updateMotto(){
    let shouldStop = false
    setTimeout(()=>{
      shouldStop = true
    },10000)
    let count = 0
    const update = ()=>{
      count ++
      if(!shouldStop){
        this.setData({
           motto:`Update count:${count}`
        },()=>{
          update()
        })
      }
    }
    update()
  }
})
