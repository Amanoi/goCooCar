const shareLocationkey = "share_location"
Page({
    data: {
        avatarURL: '',
        shareLocation: false,
    },
    async onLoad(opt) {
        console.log('unlocking cat',opt.car_id)
        const userInfo = await getApp<IAppOption>().globalData.userInfo
        this.setData({
            avatarURL: userInfo.avatarUrl,
            shareLocation: wx.getStorageSync(shareLocationkey) || false,
        })
    },
    onGetUserInfo(e: any) {
        // console.log(e.detail.value.userInfo)
        const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
        if (userInfo) {
            getApp<IAppOption>().resoveUserInfo(userInfo)
        }
    },
    onShareLocation(e: any) {
        const shareLocation: boolean = e.detail.value
        this.setData({
            shareLocation: shareLocation,
        })
        wx.setStorageSync(shareLocationkey, shareLocation)
    },
    onUnLockTap() {
        wx.getLocation({
            type: 'gcj02',
            success: loc => {
                console.log(loc, 'starting a trip', {
                    location: {},
                    //TODO:双向数据绑定
                    avatarURL: this.data.shareLocation ? this.data.avatarURL : '',
                })
                const tripID = 'trip456'

                wx.showLoading({
                    title: '解锁中',
                    mask: true,
                })
                setTimeout(() => {
                    wx.redirectTo({
                        url: `/pages/driving/driving?trip_id=${tripID}`,
                        complete: () => {
                            wx.hideLoading()
                        },
                    })
                }, 2000)
            },
            fail:()=>{
                wx.showToast({
                    icon:"none",
                    title:'请前往设置页授权位置信息',
                })
            },
        })
    }
})