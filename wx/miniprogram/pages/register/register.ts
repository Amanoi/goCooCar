import { routing } from "../../utils/routing"

Page({
    redirectURL:'',
    data: {
        licNo: '',
        name: '',
        genderIndex: 0,
        genders: ['未知', '男', '女', '其他'],
        licImgURL: '',
        birthDate: '1990-09-11',
        state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED',
    },
    onUploadLic() {
        wx.chooseImage({
            success: res => {
                if (res.tempFilePaths.length > 0) {
                    this.setData({
                        licImgURL: res.tempFilePaths[0]
                    })
                    //TODO: upload image
                    setTimeout(() => {
                        this.setData({
                            liNo: '52323123',
                            name: 'Daneil',
                            genderIndex: 1,
                            birthDate: '1989-09-07',
                        })
                    }, 1000)

                }
            }

        })
    },
    onLoad(opt:Record<'redirect',string>) {
        const o:routing.RegisterOpts = opt
        if( o.redirect){
            this.redirectURL = decodeURIComponent(o.redirect)
        }   
    },
    onGenderChange(e: any) {
        this.setData({
            genderIndex: e.detail.value,
        })
    },
    onBirthDateChange(e: any) {
        this.setData({
            birthDate: e.detail.value
        })
    },
    onSubmit() {
        //TODO: submit the form to server
        this.setData({
            state: 'PENDING',
        })
        setTimeout(() => {
            this.onLicVerified()
        }, 3000)
    },
    onResubmit() {
        this.setData({
            state: 'UNSUBMITTED',
            licImgURL: '',
        })
    },
    onLicVerified() {
        this.setData({
            state: 'VERIFIED',
        })
        if (this.redirectURL) {
            wx.redirectTo({
                url: this.redirectURL
            })
        }
        
    },
})