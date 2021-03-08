const cenPerSec = 0.7

function formatDuration(sec:number){
    const padString = (n:number)=>n<10?'0'+n.toFixed(0):n.toFixed(0)  
    const h = Math.floor(sec/3600)
    sec -= 3600*h
    const m = Math.floor(sec/60)
    sec -= 60*m
    const s = Math.floor(sec)
    return `${padString(h)}:${padString(m)}:${padString(s)}`
}
function formatFee(cents:number){
    return( cents/100).toFixed(2)
}
Page({
    timer: undefined as number|undefined,
    data: {
        elapsed: '00:00:00',
        fee: '0.00',
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
            latitude: 31.150454494018555,
            longitude: 121.51950073242188,
        },
        scale:16,
    },
    onLoad(opt) {
        console.log('current trip_id',opt.trip_id)
        this.setupLocationUpdator()
        this.setUpTimer()
    },
    onUnload(){
        wx.stopLocationUpdate() 
        this.timer && clearInterval(this.timer)
    },
    setupLocationUpdator() {
        console.log('开始监听位置变化!')
        wx.startLocationUpdate({
            fail: (res)=>console.error(res),
        })
        wx.onLocationChange((loc)=>{
            console.log(loc)
            this.setData({
                location:{
                    latitude:loc.latitude,
                    longitude:loc.longitude,
                }
            })
        })
    },
    setUpTimer(){
        let elapsedSec = 0
        let cents = 0 
        this.timer = setInterval(()=>{
            elapsedSec ++
            cents += cenPerSec
            this.setData({
                elapsed: formatDuration(elapsedSec),
                fee : formatFee(cents),
            })
        },1000)
    }
})