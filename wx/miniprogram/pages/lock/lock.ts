import {IAppOption} from "../../appoption"
import { TripService } from "../../service/trip"
import { routing } from "../../utils/routing"

const shareLocationkey = "share_location"
Page({
    carID: '',
    data: {
        avatarURL: '',
        shareLocation: false,
    },
    async onLoad(opt:Record<'car_id',string>) {
        const o:routing.LocksOpts = opt
        console.log('unlocking cat',o.car_id)
        this.carID = o.car_id
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
            getApp<IAppOption>().resolveUserInfo(userInfo)
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
            success:async loc => {
                console.log(loc, 'starting a trip', {
                    location: {},
                    //TODO:双向数据绑定
                    avatarURL: this.data.shareLocation ? this.data.avatarURL : '',
                })
                if (!this.carID){
                    console.error('no carID specified')
                    return
                }
               const trip = await TripService.CreateTrip({
                    start:loc,
                    carId:this.carID
                })
                if (!trip.id){
                    console.error('no tripID in response',trip)
                    return
                }
                wx.showLoading({
                    title: '解锁中',
                    mask: true,
                })
                setTimeout(() => {
                    wx.redirectTo({
                        url:routing.driving({
                            trip_id:trip.id,
                        }),
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