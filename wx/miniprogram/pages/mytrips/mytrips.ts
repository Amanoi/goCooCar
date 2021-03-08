import { routing } from "../../utils/routing"

Page({
    data:{
        avatarURL:'',
    },
    async onLoad() {
        const userInfo = await getApp<IAppOption>().globalData.userInfo
        this.setData({
            avatarURL: userInfo.avatarUrl,
        })
    },
    onGetUserInfo(e: any) {
        // console.log(e.detail.value.userInfo)
        const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
        if (userInfo) {
            getApp<IAppOption>().resoveUserInfo(userInfo)
        }
    },
    onRegisterTap(){
        wx.navigateTo({
            url:routing.register(),
        })
    },
})