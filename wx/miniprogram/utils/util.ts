export const formatTime = (date: Date) => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return (
    [year, month, day].map(formatNumber).join('/') +
    ' ' +
    [hour, minute, second].map(formatNumber).join(':')
  )
}

const formatNumber = (n: number) => {
  const s = n.toString()
  return s[1] ? s : '0' + s
}

//Promise 改写 getSetting
export function getSetting(): Promise<WechatMiniprogram.GetSettingSuccessCallbackResult>{
  return new Promise((reslove,reject)=>{
    wx.getSetting({
      success: reslove,
      fail: reject,
    })
  })
}

export function getUserInfo():Promise<WechatMiniprogram.GetUserInfoSuccessCallbackResult>{
  return new Promise((reslove,reject)=>{
    wx.getUserInfo({
        success:reslove,
        fail:reject,
    })
  })
}