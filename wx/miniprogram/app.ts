import { getSetting, getUserInfo } from "./utils/util"

// app.ts
App<IAppOption>({
  globalData: {
    userInfo: new Promise((resolve, reject) => {
      //应用Promise 获取用户信息
      getSetting().then(res => {
        if (res.authSetting['scope.userInfo']) {
          return getUserInfo()
        }
        //return Promise.resolve(undefined)
        return undefined
      }).then(res => {
        if (!res) {
          return
        }
        resolve(res.userInfo)
      }).catch(reject)
    }),
  },
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })

    //
  },
})