/// <reference path="./types/index.d.ts" />

interface IAppOption {
  globalData: {
    userInfo:Promise< WechatMiniprogram.UserInfo>
  }
  resoveUserInfo(userInfo:WechatMiniprogram.UserInfo): void;
}