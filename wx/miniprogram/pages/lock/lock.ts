const shareLocationkey = "share_location"
Page({
    data: {
        avatarURL: '',
        shareLocation:false,
    },
    async onLoad() {
        const userInfo = await getApp<IAppOption>().globalData.userInfo
        this.setData({
            avatarURL: userInfo.avatarUrl,
            shareLocation:wx.getStorageSync(shareLocationkey)||false,
        })
    },
    onGetUserInfo(e: any) {
        // console.log(e.detail.value.userInfo)
        const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
        if (userInfo) {
            getApp<IAppOption>().resoveUserInfo(userInfo)
        }
    },
    onShareLocation(e:any){
        const shareLocation:boolean = e.detail.value
        this.setData({
            shareLocation:shareLocation,
        })
        wx.setStorageSync(shareLocationkey, shareLocation)
    },
    onUnLockTap(){
        wx.showLoading({
            title:'解锁中',
            mask:true,
        })
        setTimeout(()=>{
            wx.redirectTo({
                url:'/pages/driving/driving',
                complete:()=>{
                    wx.hideLoading()
                },
            })
        },2000)
    },   
})