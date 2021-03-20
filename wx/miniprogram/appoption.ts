export interface IAppOption {
  globalData: {
    userInfo: Promise<WechatMiniprogram.UserInfo>
  }
  resoveUserInfo(userInfo: WechatMiniprogram.UserInfo): void;
}